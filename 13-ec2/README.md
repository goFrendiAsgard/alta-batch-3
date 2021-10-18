# Buat instance EC2

* Login ke AWS console
* Pilih menu `EC2`
* Klik `Launch Instance`
* Pilih yang free tier. Disarankan `ubuntu`
* Untuk `security group`, pastikan port 22/SSH ter ekspose
* Create key pair, dan simpan (ekstensinya `.pem`)

# Tentang Security Group

* Inbound rule: EC2 instance bisa diakses dari alamat/port apa saja?
* Outbound rule: EC2 instance bisa mengakses alamat/port sapa saja?

# Akses SSH menggunakan key pair

* Copy file `.pem` yang sudah di download ke home directory (`cp /mnt/c/Users/gofrendi/Downloads/lenovo.pem ~`)
* Ubah permission dari file `.pem` menjadi private (`chmod 700 lenovo.pem`)
* Lakukan koneksi ke ec2 (`ssh -i lenovo.pem ubuntu@ec2-54-179-44-145.ap-southeast-1.compute.amazonaws.com`)

# Install Docker

```sh
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install docker.io
sudo usermod -a -G docker ubuntu
# Kalau masih bandel: sudo chmod 777 /var/run/docker.sock
# Kalau masih bandel juga: sudo su

# download image + bikin container + run di background
docker run --name alta-nginx -p 3000:8080 bitnami/nginx -d

# jalankan container yang ada
docker start alta-nginx

# matikan container
docker stop alta-nginx
```

# Copy Paste dari/ke EC2 via ssh

```sh
# scp [-r] -i <permission.pem> [user@host:]<path> [user@host:]<path>

# copy README.md dari komputer lokal ke Ec2 di path /home/ubuntu sebagai user ubuntu dengan permission key yang ada di lenovo.pem
scp -i ~/lenovo.pem README.md ubuntu@ec2-54-179-44-145.ap-southeast-1.compute.amazonaws.com:/home/ubuntu
```