Vagrant.configure("2") do |config|

  # config.vm.box = "debian/buster64"
  
  config.vm.box = "ubuntu/jammy64"
  
  config.vm.provider "virtualbox" do |v|
      v.memory = 4096
      v.cpus = 2
  end


  config.vm.synced_folder "./.", "/home/vagrant/code"
  
end