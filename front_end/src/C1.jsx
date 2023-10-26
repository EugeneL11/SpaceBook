import {React, useState} from 'react'
function C1(props) {
    const  [number, setNum] = useState(0)
  return (
    
    <div>{props.num}</div>
  )
}

export default C1;