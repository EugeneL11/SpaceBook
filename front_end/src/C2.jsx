import {React, useState} from 'react'
function C2(props) {
    const [onC, setC] = useState(() =>{})
    const onclick = () =>{
        props.setNum(7)
    }
  return (
    
    <div onClick={onclick}>
      +
    </div>
  )
}

export default C2;