import { useDispatch, useSelector } from 'react-redux'
import './App.css'
import { increment } from './redux/counterSlice';

function App() {
  // useSelectorは状態にアクセスできる
  // state.counterはstoreのcounterから取得してきている
  let count = useSelector((state: { counter: { value: number } }) => state.counter.value);
  // dispatchはstoreに通知する
  let dispatch = useDispatch();
  return (
    <>
      <div className='APP'>
        <p>counter: {count}</p>
        <button onClick={ ()=> dispatch(increment()) }>➕</button>
      </div>
    </>
  )
}

export default App
