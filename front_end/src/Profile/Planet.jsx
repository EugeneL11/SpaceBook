import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js'
import { useLoader, useFrame } from '@react-three/fiber'
import { OrbitControls } from '@react-three/drei'
import { useRef } from 'react'

export default function Planet(props) {
    const earthRef = useRef()
    const modelpath = "./planets/" + props.planet + ".glb"

    const model = useLoader(GLTFLoader, modelpath);

    useFrame(() => {
        earthRef.current.rotation.y += 0.005;
    })

    return (
      <>
        <OrbitControls enablePan={false} enableZoom={false} rotateSpeed={0.2}/>
        <ambientLight intensity={ props.planet === "earth" ? 30 : 1 } />

        <primitive ref={earthRef} object={model.scene} scale={0.006} />
      </>
    );
}