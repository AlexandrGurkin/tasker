sudo sysctl net.ipv4.tcp_max_tw_buckets=65536
sudo sysctl net.ipv4.tcp_tw_recycle=1
sudo sysctl net.ipv4.tcp_tw_reuse=0
sudo sysctl net.ipv4.tcp_max_syn_backlog=131072
sudo sysctl net.ipv4.tcp_syn_retries=3
sudo sysctl net.ipv4.tcp_synack_retries=3
sudo sysctl net.ipv4.tcp_retries1=3
sudo sysctl net.ipv4.tcp_retries2=8
sudo sysctl net.ipv4.tcp_rmem="16384 174760 349520"
sudo sysctl net.ipv4.tcp_wmem="16384 131072 262144"
sudo sysctl net.ipv4.tcp_mem="262144 524288 1048576"
sudo sysctl net.ipv4.tcp_max_orphans=65536
sudo sysctl net.ipv4.tcp_fin_timeout=10
sudo sysctl net.ipv4.tcp_low_latency=1
sudo sysctl net.ipv4.tcp_syncookies=0
sudo sysctl net.netfilter.nf_conntrack_max=1048576