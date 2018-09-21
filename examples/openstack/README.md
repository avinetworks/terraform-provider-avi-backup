# OpenStack terraform example

## Instruction to lanch an instance 

To run OpenStack example you should know the following things of OpenStack
```
 - Openstack credentials
 - Network name
 - Subnet 
 - You should have image file 
 - Controller default password
```

## Commands to run terraform scripts 
```
terraform init
terraform plan
terraform apply
```

To launch an instance just go inside the launch_instance directory and run following commands while running just pass the required arguments that it prompts
```
terraform init
terraform plan
terraform apply
```

To set password to the controller go inside the avi_user_account and run same commands

To set up OpenStack cloud on controller go inside the setup_cloud directory and run the same commands

To create an application on controller go inside the web_app directory and run the same commands 


