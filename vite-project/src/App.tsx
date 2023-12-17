import { useSelector } from 'react-redux'
import './App.css'

function App() {
  // useSelectorは状態にアクセスできる
  // state.counterはstoreのcounterから取得してきている
  let count = useSelector((state: { counter: { value: number } }) => state.counter.value);

  return (
    <>
      <div className='APP'>
        <p>counter: {count}</p>
      </div>
    </>
  )
}

export default App
