# Langkah-langkah

* Buat akun okteto
* Download kube-config
* Update `~/.kube/config` atau `C:\\Users\<user-name>\.kube\config`
* Pastikan `kubectl config get-contexts` mengarah ke okteto, jika tidak lakukan `kubectl config use-context <context-name>`
* Deploy mysql (`./mysql/deploy.sh`)
* Deploy go app (`./go-app/deploy.sh`)


# Catatan

Vendor lain, seperti AWS dan GCP memiliki mekanisme sendiri untuk update kubeconfig.