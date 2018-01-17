private_ip = '192.168.33.10'

start_services = <<-EOF
set -ex

systemctl start zookeeper.service
/usr/share/zookeeper/bin/zkCli.sh create /mesos
/usr/share/zookeeper/bin/zkCli.sh create /marathon
systemctl restart mesos_executors.slice
systemctl restart mesos-master.service
systemctl restart mesos-slave.service
systemctl restart marathon.service
systemctl restart zookeeper.service
EOF

install_script = <<-EOF
apt-key adv --keyserver keyserver.ubuntu.com --recv E56151BF
echo "deb http://repos.mesosphere.io/ubuntu xenial main" |
  tee /etc/apt/sources.list.d/mesosphere.list
apt-get update
apt-get install -y mesos marathon chronos vim openjdk-9-jre-headless
apt-get install -y curl apt-transport-https software-properties-common
url -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository 'deb [arch=amd64] https://download.docker.com/linux/ubuntu xenial stable'
apt-get update
apt-get install -y --allow-unauthenticated docker-ce
mkdir -p /usr/lib/jvm/java-9-openjdk-amd64/conf/management/
ln -s /etc/java-9-openjdk/management/management.properties /usr/lib/jvm/java-9-openjdk-amd64/conf/management/management.properties
echo "#{private_ip}" | tee /etc/mesos-master/ip
echo "#{private_ip}" | tee /etc/mesos-slave/ip
echo "#{private_ip}" | tee /etc/mesos-master/hostname
echo "1" | tee /etc/zookeeper/conf/myid
echo "mesops-def" | tee /etc/mesos-master/cluster
echo "docker,mesos" | tee /etc/mesos-slave/containerizers
printf "MARATHON_MASTER=192.168.33.10:5050\nZK=zk://127.0.0.1:2181/marathon" | tee /etc/default/marathon
echo "#{start_services}" | tee /usr/local/bin/start-services.sh
chmod +x /usr/local/bin/start-services.sh
EOF

Vagrant.configure('2') do |config|
  config.vm.box      = 'ubuntu/xenial64'
  config.vm.hostname = 'mesos'

  config.vm.network :private_network, ip: private_ip
  # for mesos web UI.
  config.vm.network :forwarded_port, guest: 5050, host: 5050
  # for Marathon web UI
  config.vm.network :forwarded_port, guest: 8080, host: 8080

  config.vm.provision 'shell', inline: install_script

  config.vm.provision 'shell', inline: '/usr/local/bin/start-services.sh'

  config.vm.provider 'virtualbox' do |vb|
    vb.customize ['modifyvm', :id, '--memory', "#{1024*2}"]
    vb.customize ["modifyvm", :id,  '--cpus',  '2']
  end
end
