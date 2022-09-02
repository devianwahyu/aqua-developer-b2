# **Aqua Developer Batch 2 Day 10**

Tanggal: 2 September 2022

Materi: OneFish Design System

## **Pengertian Design System**

Standarisasi desain reusable yang dapat digunakan dalam mengembangkan frontend/tampilan aplikasi dalam bentuk component

why menggunakan design system?
- Lebih cepat dalam pengembangan dan melakukan scale
- Menjadikan tampilan lebih konsisten
- Dapat jadi referensi bagi junior engineer

## **Pengertian OneFish Design System**

OneFish merupakan design system yang dipakai oleh internal eFishery dalam mengembangkan aplikasi perusahaan. Design system ini mengambil dari ChakraUI dan berbasis ReactJS.

## **ReactJS Basic**

ReactJS merupakan library Javascript yang digunakan dalam pengembangan UI suatu aplikasi. Salah satu cara instalasinya adalah dengan ```create-react-app```

requirements:
- NodeJS

instalasi
```npx create-react-app app-name```

## **Instalasi**

instalasi OneFish Design System:
```
- install react app
  npx create-react-app my-app

- install ChakraUI
  npm i @chakra-ui/react @emotion/react @emotion/styled framer-motion

- post install OneFish
  npm config set //balong.efishery.com/:_authToken="<token>"
  npm config set registry=https://balong.efishery.com
  npm config set always-auth=true

- install
  yarn add @efishery/onefish --registry https://balong.efishery.com
  yarn add @chakra-ui/react@^1.8.8 @chakra-ui/icons @chakra-ui/system @emotion/react@^11 @emotion/styled@^11 framer-motion@^4.1.17 @fontsource/poppins
```

## **Cara Penggunaan**
Tidak banyak yang bisa diimplementasikan, karena untuk menggunakan OneFish Design System harus menggunakan token internal eFishery. Tapi secara general cara penggunaanya mirip dengan penggunaan ChakraUI. Apabila ada design system yang belum terimplementasikan ke OneFish maka developer dapat menggunakan ChakraUI dengan properti menyesuaikan dengan design system eFishery.