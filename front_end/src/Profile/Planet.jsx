import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js'
import { useFrame, useThree } from '@react-three/fiber'
import { OrbitControls, Html } from '@react-three/drei'
import { useRef, useState, useEffect } from 'react'

export default function Planet(props) {
    const earthRef = useRef()
    const modelpath = "./planets/" + props.planet + ".glb"

    const [model, setModel] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const loader = new GLTFLoader();
        loader.load(modelpath, (gltf) => {
            setModel(gltf);
            setLoading(false);
        });
    }, [modelpath]);

    // Create a Map
    const lightMap = new Map([
      ["mercury", 10],
      ["venus", 1],
      ["earth", 30],
      ["mars", 1.5],
      ["jupiter", 1.5],
      ["saturn", 1],
      ["uranus", 1],
      ["neptune", 1],
      ["pluto", 5]
    ]);

    useFrame(() => {
        if (earthRef.current) {
            earthRef.current.rotation.y += 0.005;
        }
    })

    return (
      <>
        <OrbitControls enablePan={false} enableZoom={false} rotateSpeed={0.2}/>
        <ambientLight intensity={ lightMap.get(props.planet) } />

        {model && <primitive ref={earthRef} object={model.scene} scale={
          props.planet == "saturn" ? 0.0025 : 0.006
        }/>}

        {loading && (
          <Html center>
            <div className=' text-white text-xl text-center w-20 font-bold'>
              Loading 3D model...
            </div>
          </Html>
        )}
      </>
    );
}
