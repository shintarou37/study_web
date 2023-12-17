import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { Provider } from 'react-redux'
import { store } from './redux/store'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    {/* APPコンポーネントであればstoreを共有で使用できるように設定する */}
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>,
)
