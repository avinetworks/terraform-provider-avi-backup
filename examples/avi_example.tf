provider "aws" {
  access_key = "ABC"
  secret_key = "XYZ"
  region     = "us-west-2"
}

variable "aws_vpc_id" {
  type = "string"
  default = "vpc-19295f7c"
}

variable "avi_admin_user" {
  type = "string"
  default = "admin"
}

variable "avi_admin_password" {
  type = "string"
  default = "admin"
}

/*resource "aws_instance" "avi_controller" {
  ami = "ami-5f39c527"
  instance_type = "t2.large"
  subnet_id = "${aws_subnet.grastogi-subnet.id}"
  tags {
    Name = "grastogi-terraform-controller"
  }
}

output "aws_controller_instance" {
  value = "${aws_instance.avi_controller.tags.Name}"
}

output "avi_controller_ip" {
  value = "${aws_instance.avi_controller.private_ip}"
}


resource "aws_subnet" "grastogi-subnet" {
  vpc_id = "${var.aws_vpc_id}"
  cidr_block = "10.144.42.0/24"
  tags {
    Name = "grastogi-subnet"
  }
}

output "aws_subnet" {
  value = "${aws_subnet.grastogi-subnet.tags.Name}"
}
*/

provider "avi" {
  avi_username = "admin"
  //avi_password = "admin"
  avi_password = "avi123"
  avi_controller = "10.10.25.42"
  //avi_controller = "${aws_instance.avi_controller.private_ip}"
  avi_tenant = "admin"
}



resource "avi_pool" "testpool" {
  obj= {
    name= "p2",
    /*
    health_monitor_refs= ["api/healthmonitor/?name=System-HTTP"]
    servers= [
      {
        ip= {
          type= "V4",
          addr= "10.90.64.66",
        },
      port= 9000
      },
    ],
    */
  }
}


output "pool_uuid" {
  value = "${avi_pool.testpool.id}"
}
