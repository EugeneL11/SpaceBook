import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js'
import { useLoader, useFrame } from '@react-three/fiber'
import { OrbitControls } from '@react-three/drei'
import { useRef } from 'react'

export default function Planet() {
    const earthRef = useRef()
    const model = useLoader(GLTFLoader, "./earth.glb");

    useFrame(() => {
        earthRef.current.rotation.y += 0.005;
    })

    return (
      <>
        <OrbitControls enablePan={false} enableZoom={false} rotateSpeed={0.2}/>
        <ambientLight intensity={30} />

        <primitive ref={earthRef} object={model.scene}  scale={0.006} />
      </>
    );
}