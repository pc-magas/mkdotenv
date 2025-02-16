Vagrant.configure("2") do |config|

  # config.vm.box = "debian/buster64"
  
  config.vm.box = "fedora/39-cloud-base"
  
  config.vm.provider "virtualbox" do |v|
      v.memory = 4096
      v.cpus = 2
  end


  config.vm.synced_folder "./.", "/home/vagrant/code"
  config.vm.synced_folder "~/.gnupg", "/home/vagrant/.gnupg"
  
  # config.vm.provision "shell", privileged: true, inline: <<-SHELL
  #   apt-get update
  #   DEBIAN_FRONTEND=noninteractive apt-get dist-upgrade -y
  #   DEBIAN_FRONTEND=noninteractive apt-get install -y golang-1.23-go debhelper make sbuild devscripts
  #   usermod -aG sbuild vagrant
  # SHELL

end