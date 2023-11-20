import { React, useState } from "react";
import static1 from "../Static.js";
function Register(props) {
    const toggleLogin = props.toggleLogin
    const toggleHomepage = props.toggleHomepage

    const [fullName, setName] = useState("")
    const [email, setEmail] = useState("")
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setError] = useState(null)
    const registerAction = () => {
        //ask backend
        const localhost = "http://localhost:8080"
        
        const path = `/register/${encodeURIComponent(email)}/${encodeURIComponent(password)}/${encodeURIComponent(fullName)}/${encodeURIComponent(username)}`
        fetch(`${localhost}${path}`).then(res => { res.json().then(data =>{
            if(data.error === "email already in use"){
                setError(data.error)
            }
            else if(data.error ==="unable to create account at this time"){
                setError(data.error)
            }
            else if(data.error === "user name not availible"){
                setError(data.error)
            }
            else{
                static1.userID = data.id
                console.log(static1.userID)
                toggleHomepage()
            }
        })})
        /*
        const path = "/testInsert/val"
        fetch(`${localhost}${path}`).then(res => {res.json().then(data =>{
            console.log(data)
        })})*/
    }

    return (
    <div class="flex flex-col items-center">
        <h1 class="text-6xl py-11">SpaceBook</h1>

        <div class="bg-white text-black text-center text-lg p-6 sm:px-12 rounded-md w-3/4 sm:w-1/2 lg:w-2/5 xl:w-1/3 min-w-fit">
            <h4 class="text-3xl mb-2">Welcome!</h4>
            <div class="w-full">
                <label for="username" class="block text-left">Full Name</label>
                <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter Full Name" type="text" value={fullName} onChange={e => setName(e.target.value)}></input>
            </div>
            <div class="w-full mt-4">
                <label for="password" class="block text-left">Email Address</label>
                <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter Email Address" type="text" value={email} onChange={e => setEmail(e.target.value)}></input>
            </div>
            <div class="w-full mt-4">
                <label for="username" class="block text-left">Username</label>
                <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter Username" type="text" value={username} onChange={e => setUsername(e.target.value)}></input>
            </div>
            <div class="w-full mt-4">
                <label for="password" class="block text-left">Password</label>
                <input class="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" placeholder="Enter Password" type="password" value={password} onChange={e => setPassword(e.target.value)}></input>
            </div>
            <button class="bg-purple-200 px-7 py-3 mt-6 rounded-lg" onClick={registerAction}>Sign Up</button>
        </div>

        <button class="text-xl mt-12" onClick={toggleLogin}>Log In</button>
        <h1>{errorMessage}</h1>
    </div>
    );

}

export default Register;