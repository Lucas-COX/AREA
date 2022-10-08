import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { ToastContainer } from 'react-toastify'

import 'react-toastify/dist/ReactToastify.css';


function MyApp({ Component, pageProps }: AppProps) {
  return (
    <div>
      <Component {...pageProps} />
      <ToastContainer autoClose={4000} />
    </div>
  )
}

export default MyApp
