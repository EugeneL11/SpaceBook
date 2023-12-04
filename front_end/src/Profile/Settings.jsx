import { React, useState } from "react";
import backPic from '../images/back.png';
import axios, { formToJSON } from 'axios'
import currentUser from "../Static";
import { serverpath } from "../Path";
function Settings(props) {
    const toggleLogin = props.toggleLogin
    const toggleMyProfile = props.toggleMyProfile
    const [fullName, setFullName] = useState(currentUser.full_name)
    const [bio, setBio] = useState(currentUser.bio)
    const [planet, setPlanet] = useState(currentUser.planet)
    const [image, setImage] = useState(null)
    const handleFullName = (event) => {
        setFullName(event.target.value)
    }

    const handleBio = (event) => {
        setBio(event.target.value)
    }

    const handlePlanet = (event) => {
        setPlanet(event.target.value)
    }
    const setFile = (e) =>{
        setImage(e.target.files[0])
    }
    async function updateSettings() {
        // if (bio === "") {
        //     setBio(" ")
        // }
        // if (planet === "") {
        //     setPlanet()
        // }
        const path = `/updateuserprofile/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(fullName)}/${encodeURIComponent(planet)}/${encodeURIComponent(bio)}`
        const res = await axios.put(`${serverpath}${path}`)
        const data = res.data
        if (data.status !== "no error") {
            console.log(data.status)
            return;
        }
        if(image !== null){
            const setimagepath = `/uploadprofileimage/${encodeURIComponent(currentUser.userID)}`
            const formData = new FormData();
            formData.append("image", image);
            const imageres = await axios.post(`${serverpath}${setimagepath}`, formData)
            console.log(imageres)
        }
        currentUser.planet = planet
        currentUser.full_name = fullName
        currentUser.bio = bio
        toggleMyProfile()
    }
    const deleteAcc = () => {
        const path = `/deleteuser/${encodeURIComponent(currentUser.userID)}`
        axios.delete(`${serverpath}${path}`).then(res => {
            if (res.data.status == "no error"){
                toggleLogin()
            }
            else{
                console.log(res.data.status)
            }
        })
    }
    return (
    <div className="flex flex-col items-center">
        <button className="mb-5 w-fit ml-6 mr-auto text-3xl hover:text-purple-300" onClick={toggleMyProfile}>{'←'}</button>
        
        <div className="flex flex-col bg-white text-black text-start text-lg mt-[-20px] mb-10 md:py-6 sm:px-16 lg:px-24 p-6 rounded-xl w-full sm:w-3/4 lg:w-1/2 min-w-fit">
            <h1 className="text-center text-3xl pt-2">Settings</h1>
            <label className="mt-4">Change Full Name: </label>
            <input 
                className="w-full border-b-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" 
                placeholder="Enter New Full Name" type="text"
                value={fullName} onChange={handleFullName}>
            </input>
        
            <div className="mt-4">Change Profile Picture: </div>
            <input type="file" className="form-input text-sm" onChange={setFile}></input>

            <div className="mt-4">Edit Bio:</div>
            <textarea 
                className="form-textarea border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0" 
                rows="3" placeholder=" Edit your bio..."
                value={bio} onChange={handleBio}>
            </textarea>

            <div className="mt-4">Change Home Planet: </div>
            <select 
                className="form-select border-2 border-gray-700 focus:outline-none focus:border-gray-300 focus:ring-0"
                value={planet} onChange={handlePlanet}>
                <option value="Mercury">Mercury</option>
                <option value="Venus">Venus</option>
                <option value="Earth">Earth</option>
                <option value="Mars">Mars</option>
                <option value="Jupiter">Jupiter</option>
                <option value="Saturn">Saturn</option>
                <option value="Uranus">Uranus</option>
                <option value="Neptune">Neptune</option>
                <option value="Pluto">Pluto</option>
            </select>
            
            <button className="bg-purple-300 hover:bg-purple-400 px-4 py-2 mt-5 w-fit self-center rounded-lg" onClick={updateSettings}>Apply Changes</button>

            <button className="w-fit self-center mt-6 hover:text-blue-300" onClick={toggleLogin}>Log Out</button>
            <button className="w-fit self-center hover:text-red-600" onClick={deleteAcc}>Delete Account</button>
        </div>
    </div>
    );

}

export default Settings;