### rid-apply

#database schema

![image](https://user-images.githubusercontent.com/22321111/120995348-85078980-c7af-11eb-9aa4-5c36b1e00a6e.png)

#alur aplikasi
1. user melalui client/front-end login menggunakan google.
2. google redirect ke back-end.
3. back-end verifikasi state.
4. back-end menukarkan code menjadi access token.
5. access token digunakan untuk mengambil data user.
![image](https://user-images.githubusercontent.com/22321111/120996740-cc424a00-c7b0-11eb-9dbd-53b0f9c4bf86.png)

6. access token dikembalikan ke client/front-end sebagai response.
![image](https://user-images.githubusercontent.com/22321111/120996641-b3399900-c7b0-11eb-9907-0ec74fda345c.png)

7. access token bisa digunakan untuk authorization ke back-end dan di-refresh jika expired
![image](https://user-images.githubusercontent.com/22321111/120997603-8c2f9700-c7b1-11eb-8ca8-fdd2e8458079.png)
