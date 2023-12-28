import { useDispatch, useSelector } from 'react-redux'
import '../App.css'
import { decrement, increment, incrementByInput } from '../redux/counterSlice';
import { useState } from 'react';

function Redux() {
  // useSelectorは状態にアクセスできる
  // state.counterはstoreのcounterから取得してきている
  const count = useSelector((state: { counter: { value: number } }) => state.counter.value);
  // dispatchはstoreに通知する
  const dispatch = useDispatch();
  const [inputNum, setInputNum] = useState("1");
  return (
    <>
      <div className='APP'>
        <p>counter: {count}</p>
        <input onChange={ (e)=> setInputNum(e.target.value)}></input>
        <button onClick={ ()=> dispatch(increment()) }>➕</button>
        <button onClick={ ()=> dispatch(decrement()) }>-</button>
        {/* incrementByInputaction.payloadにinputNumの値が入る */}
        <button onClick={ ()=> dispatch(incrementByInput(Number(inputNum)))}>追加</button>
      </div>
    </>
  )
}

export default Redux