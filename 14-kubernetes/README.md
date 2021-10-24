# Langkah-langkah

* Buat akun okteto
* Download kube-config
* Update `~/.kube/config` atau `C:\\Users\<user-name>\.kube\config`
* Pastikan `kubectl config get-contexts` mengarah ke okteto, jika tidak lakukan `kubectl config use-context <context-name>`
* Deploy mysql (`./mysql/deploy.sh`)
* Deploy go app (`./go-app/deploy.sh`)


# Catatan

* Cloud provider lain, seperti AWS dan GCP memiliki mekanisme sendiri untuk update kubeconfig. Biasanya harus menginstall CLI tool yang mereka sediakan
* Untuk MySQL sebaiknya menggunakan `statefulSet`, bukan `deployment`, namun dalam contoh ini kita menggunakan `deployment` saja, karena okteto free plan tidak mengijinkan akses `volume` API
* Supaya service bisa diakses dari luar cluster, kita perlu menggunakan `ingress`