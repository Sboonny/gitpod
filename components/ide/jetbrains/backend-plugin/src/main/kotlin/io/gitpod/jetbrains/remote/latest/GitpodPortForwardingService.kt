// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package io.gitpod.jetbrains.remote.latest

import com.intellij.openapi.diagnostic.thisLogger
import com.intellij.openapi.project.Project
import com.intellij.remoteDev.util.onTerminationOrNow
import com.intellij.util.application
import com.jetbrains.codeWithMe.model.RdPortType
import com.jetbrains.rd.platform.util.lifetime
import com.jetbrains.rd.util.lifetime.LifetimeStatus
import com.jetbrains.rdserver.portForwarding.ForwardedPortInfo
import com.jetbrains.rdserver.portForwarding.PortForwardingManager
import com.jetbrains.rdserver.portForwarding.remoteDev.PortEventsProcessor
import io.gitpod.jetbrains.remote.GitpodManager
import io.gitpod.supervisor.api.Status
import io.gitpod.supervisor.api.StatusServiceGrpc
import io.grpc.stub.ClientCallStreamObserver
import io.grpc.stub.ClientResponseObserver
import io.ktor.utils.io.*
import java.util.concurrent.CompletableFuture
import java.util.concurrent.TimeUnit

@Suppress("UnstableApiUsage")
class GitpodPortForwardingService(private val project: Project) {
    companion object {
        const val FORWARDED_PORT_LABEL = "gitpod"
    }

    init { start() }

    private fun start() {
        if (application.isHeadlessEnvironment) return

        observePortsListWhileProjectIsOpen()
    }

    private fun observePortsListWhileProjectIsOpen() = application.executeOnPooledThread {
        while (project.lifetime.status == LifetimeStatus.Alive) {
            try {
                observePortsList().get()
            } catch (throwable: Throwable) {
                when (throwable) {
                    is InterruptedException, is CancellationException -> break
                    else -> thisLogger().error(
                            "gitpod: Got an error while trying to get ports list from Supervisor. " +
                                    "Will try again in 1 second.",
                            throwable
                    )
                }
            }

            TimeUnit.SECONDS.sleep(1)
        }
    }

    private fun observePortsList(): CompletableFuture<Void> {
        val completableFuture = CompletableFuture<Void>()

        val statusServiceStub = StatusServiceGrpc.newStub(GitpodManager.supervisorChannel)

        val portsStatusRequest = Status.PortsStatusRequest.newBuilder().setObserve(true).build()

        val portsStatusResponseObserver = object :
                ClientResponseObserver<Status.PortsStatusRequest, Status.PortsStatusResponse> {
            override fun beforeStart(request: ClientCallStreamObserver<Status.PortsStatusRequest>) {
                project.lifetime.onTerminationOrNow { request.cancel("gitpod: Project terminated.", null) }
            }
            override fun onNext(response: Status.PortsStatusResponse) { handlePortStatusResponse(response) }
            override fun onCompleted() { completableFuture.complete(null) }
            override fun onError(throwable: Throwable) { completableFuture.completeExceptionally(throwable) }
        }

        statusServiceStub.portsStatus(portsStatusRequest, portsStatusResponseObserver)

        return completableFuture
    }

    private fun handlePortStatusResponse(response: Status.PortsStatusResponse) {
        val portNumberToServedStatusMap = mutableMapOf<Int, Boolean>()

        for (port in response.portsList) {
            val isServed = (port.exposed != null || port.tunneled != null) && port.served

            portNumberToServedStatusMap[port.localPort] = isServed
        }

        updateForwardedPortsList(portNumberToServedStatusMap)
    }

    private fun updateForwardedPortsList(portNumberToServedStatusMap: Map<Int, Boolean>) {
        val portForwardingManager = PortForwardingManager.getInstance(project)
        val forwardedPortsList = portForwardingManager.getForwardedPortsWithLabel(FORWARDED_PORT_LABEL)

        portNumberToServedStatusMap.forEach { (hostPort, isServed) ->
            if (isServed && !forwardedPortsList.containsKey(hostPort)) {
                val portEventsProcessor = object : PortEventsProcessor {
                    override fun onPortForwarded(hostPort: Int, clientPort: Int) {
                        thisLogger().info("gitpod: Forwarded port $hostPort to client's port $clientPort.")
                    }

                    override fun onPortForwardingFailed(hostPort: Int, reason: String) {
                        thisLogger().error("gitpod: Failed to forward port $hostPort: $reason")
                    }
                }

                val portInfo = ForwardedPortInfo(
                        hostPort,
                        RdPortType.HTTP,
                        FORWARDED_PORT_LABEL,
                        emptyList(),
                        portEventsProcessor
                )

                portForwardingManager.forwardPort(portInfo)
            }

            if (!isServed && forwardedPortsList.containsKey(hostPort)) {
                portForwardingManager.removePort(hostPort)
                thisLogger().info("gitpod: Stopped forwarding port $hostPort.")
            }
        }
    }
}
