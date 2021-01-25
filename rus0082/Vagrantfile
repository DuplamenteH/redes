Vagrant.configure(2) do |config|
  #Ubuntu 20.10 64 bit
  config.vm.box = "ubuntu/groovy64"
  #Nome da máquina virtual
  config.vm.hostname = "rus0082"
  #Permite integração com console/terminal/linha de comando
  config.ssh.forward_x11 = true
  #config.vm.network "forwarded_port", guest: 8888, host: 8888

  #Configurando a máquina virtual com as dependências
  config.vm.provision "shell", inline: <<-SHELL
     #Atualizando ubuntu
     sudo apt update
     sudo apt upgrade -y
     sudo apt autoremove
     #Dando permissões ao scripts bash
     find /vagrant -name "*.sh" | xargs chmod -v 744
     #Se máquina virtual rodando no Windows, converter quebras de linhas ao estilo UNIX
     sudo apt-get install -y dos2unix
     printf "Usando dos2unix para converter arquivos ao formato Unix se necessário..."
     find /vagrant -name "*" -type f | xargs dos2unix -q
     #Instalação make
     sudo apt install make
     #Download go
     cd /tmp/
     wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
     #Extraindo go
     tar -xvf go1.15.6.linux-amd64.tar.gz
     sudo mv go /usr/local/
     #Configurar variável de ambiente para GO
     echo "export GOPATH=/vagrant/" >> /home/vagrant/.bashrc
     echo "export PATH=/usr/local/go/bin:$PATH" >> /home/vagrant/.bashrc
     source ~/.profile
     #Inicia em /vagrant, que é um link para diretório do trabalho na máquina real
     if ! grep -Fxq "cd /vagrant" /home/vagrant/.bashrc
     then
      echo "cd /vagrant" >> /home/vagrant/.bashrc
     fi
  SHELL

  #Configurações de memória e CPU
  config.vm.provider "virtualbox" do |vb|
    vb.customize ["modifyvm", :id, "--cpuexecutioncap", "100"]
    vb.memory = 2048
    vb.cpus = 1
  end

end
