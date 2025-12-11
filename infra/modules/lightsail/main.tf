resource "aws_lightsail_static_ip" "app_ip" {
  name = "${var.instance_name}-ip"
}

resource "aws_lightsail_disk" "rss2go_db" {
  name              = "${var.instance_name}-db"
  size_in_gb        = var.disk_size_gb
  availability_zone = "${var.aws_region}a"
}

resource "aws_lightsail_instance" "app" {
  name              = var.instance_name
  availability_zone = "${var.aws_region}a"
  blueprint_id      = "docker"
  bundle_id         = var.instance_plan

  user_data = <<-EOF
    #!/bin/bash
    sudo yum update -y
    sudo amazon-linux-extras enable docker
    sudo yum install -y docker
    sudo service docker start
    sudo usermod -a -G docker ec2-user

    # mount persistent SQLite disk
    sudo mkdir -p /mnt/rss2go-db
    sudo mount /dev/xvdf /mnt/rss2go-db
    sudo chown -R ec2-user:ec2-user /mnt/rss2go-db

    # pull container
    docker pull ${var.container_image}

    # run container
    docker run -d --name rss2go \
      -p 8080:8080 \
      -v /mnt/rss2go-db:/app/internal/database \
      -e AUTH=${var.enable_auth} \
      -e APP_USER=${var.app_user} \
      -e APP_PASS=${var.app_pass} \
      ${var.container_image}
  EOF
}

resource "aws_lightsail_static_ip_attachment" "ip_attach" {
  static_ip_name = aws_lightsail_static_ip.app_ip.name
  instance_name  = aws_lightsail_instance.app.name
}

resource "aws_lightsail_disk_attachment" "db_attach" {
  disk_name     = aws_lightsail_disk.rss2go_db.name
  instance_name = aws_lightsail_instance.app.name
  disk_path     = "/dev/xvdf"
}
