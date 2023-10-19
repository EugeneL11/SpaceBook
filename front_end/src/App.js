import logo from './logo.svg';
import './App.css';
import FetchTest from './FetchTest'
import C1 from './C1'
import C2 from './C2'
import { useState } from 'react';
function App() {
  const [num, setNum] = useState(0)
  return (
    /*<div>
    <h1 className="text-3xl font-bold underline m-10">
      Hey guys! This was styled with Tailwind CSS!
      <br /> <br />
      - Omar
    </h1>
    <FetchTest/>
    </div>*/
    <div> <C1 num = {num}></C1>
    <C2 num = {num} setNum = {setNum}/></div>
   
  )
}

export default App;
