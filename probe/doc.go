/*
延迟探测模块

最初的计划是使用原始套接字实现，但发现原始套接字的各种特性实在很难琢磨。这里才用pcap
截获目标机的返回包。主要是ICMP包和TCP的ACK+SYN包。

*/
package probe
