### rid-apply

#database schema
![image](https://user-images.githubusercontent.com/22321111/120995348-85078980-c7af-11eb-9aa4-5c36b1e00a6e.png)

#alur aplikasi
1. user melalui client/front-end login menggunakan google.
2. google redirect ke back-end.
3. back-end verifikasi state.
4. back-end menukarkan code menjadi access token.
5. access token digunakan untuk mengambil data user.
6. access token dikembalikan ke client/front-end sebagai response.
