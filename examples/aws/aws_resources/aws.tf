provider "aws" {
  //access_key = "${var.aws_access_key}"
  //secret_key = "${var.aws_secret_key}"
  shared_credentials_file = "${var.aws_creds_file}"
  region     = "${var.aws_region}"
}


//data "aws_iam_instance_profile" "avi_controller_iam" {
//  name = "AviControllerRole"
//}

resource "aws_instance" "avi_controller" {
  ami           = "${var.avi_controller_ami}"
  instance_type = "c4.2xlarge"
  key_name = "no_def_pass_ctrl"
//  subnet_id     = "${aws_subnet.terraform-subnet.0.id}"
  subnet_id = "subnet-d6566c8c"
  security_groups = ["sg-0fb950f781cfb7543"]
//  iam_instance_profile = "${data.aws_iam_instance_profile.avi_controller_iam.name}"

  tags {
    Name = "${var.project_name}-terraform-controller"
  }
  provisioner "remote-exec" {
    connection={
      user="admin"
      private_key="${file("/home/shrikantc/Downloads/no_def_pass_ctrl.pem")}"
    }
    inline = [
      "echo -e ${aws_instance.avi_controller.public_ip}\\n${var.avi_password}|sudo /opt/avi/scripts/initialize_admin_user.py",
    ]
  }
//  user_data = "echo -e 34.219.106.121\\navi123$%^|sudo /opt/avi/scripts/initialize_admin_user.py"
}
//
//resource "aws_instance" "avi_controller_2" {
//  ami           = "${var.avi_controller_ami}"
//  instance_type = "c4.2xlarge"
//  subnet_id     = "${aws_subnet.terraform-subnet.0.id}"
//  iam_instance_profile = "${data.aws_iam_instance_profile.avi_controller_iam.name}"
//
//  tags {
//    Name = "${var.project_name}-terraform-controller-2"
//  }
//}
//
//resource "aws_instance" "avi_controller_3" {
//  ami           = "${var.avi_controller_ami}"
//  instance_type = "c4.2xlarge"
//  subnet_id     = "${aws_subnet.terraform-subnet.0.id}"
//  iam_instance_profile = "${data.aws_iam_instance_profile.avi_controller_iam.name}"
//
//  tags {
//    Name = "${var.project_name}-terraform-controller-3"
//  }
//}


output "aws_controller_instance" {
  value = "${aws_instance.avi_controller.tags.Name}"
}

output "avi_controller_ip" {
  value = "${aws_instance.avi_controller.private_ip}"
}

output "avi_controller_public_ip" {
  value = "${aws_instance.avi_controller.public_ip}"
}
//output "avi_controller_2_ip" {
//  value = "${aws_instance.avi_controller_2.private_ip}"
//}
//
//output "avi_controller_3_ip" {
//  value = "${aws_instance.avi_controller_3.private_ip}"
//}

//output "avi_controller_iam_role" {
//  value = "${data.aws_iam_instance_profile.avi_controller_iam.name}"
//}

//resource "aws_subnet" "terraform-subnet" {
//  count = "${length(var.aws_subnets)}"
//
//  vpc_id            = "${var.aws_vpc_id}"
//  availability_zone = "${var.aws_availability_zones[count.index]}"
//  cidr_block        = "${var.aws_subnets[count.index]}/${var.aws_subnet_mask}"
//  tags {
//    Name    = "${var.project_name}-terraform-subnet-${count.index}"
//    Project = "${var.project_name}-terraform-subnets"
//  }
//}

//output "aws_subnets" {
//  value = "${aws_subnet.terraform-subnet.*.tags.Name}"
//}

