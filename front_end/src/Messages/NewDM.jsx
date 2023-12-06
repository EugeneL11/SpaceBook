import { React, useState , useEffect } from "react";
import currentUser from "../Static.js";
import {serverpath} from "../Path.js";
import axios from 'axios'
function Person(props) {
    const recieverID = props.userID
    const user_pic_url = props.user_pic_url
    const toggleDMMessage = props.toggleDMMessage
    //clicking the profile should toggle to the dm message page with that user, while also posting the following
    // const path = `/newdm/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(props.username)}`
    // axios.get(`${serverpath}${path}`).then((res) => {
    //     const data = res.data
    //     console.log(data)
    // })

    const handleNewDM = () => {
        const path = `/newdm/${currentUser.userID}/${recieverID}`
        axios.post(`${serverpath}${path}`).then((res) => {
            const data = res.data
            if (data.status !== "no error") {
                console.log(data.status)
            }
        })
        toggleDMMessage(props.userID, props.username)
    }
    return (
        <div onClick={ handleNewDM }
        
        className="flex items-center w-11/12 sm:w-3/4 lg:w-1/2 min-w-fit bg-blue-500 space-x-4 rounded-md hover:cursor-pointer hover:bg-blue-300">

            <img 
                src={props.user_pic_url}
                alt={props.username} 
                className="w-12 h-12 rounded-full aspect-square p-2"
            ></img>
            <p className="text-lg">{props.username}</p>
            {/* <button onClick={() => {toggleOtherProfile(props.username, toggleSearchUser)}}> See Other Profile: {person}</button> */}
        </div>
    );
}

function NewDM(props) {
    const samplePeople = [
        {username: "Vic", user_pic_url: "./jupiter.jpg"},
        {username: "Kevin", user_pic_url: "./jupiter.jpg"}
    ]

    const toggleHomepage = props.toggleHomepage
    const toggleDMMessage = props.toggleDMMessage
    const toggleDMList = props.toggleDMList
    const toggleOtherProfile = props.toggleOtherProfile
    const toggleSearchUser = props.toggleSearchUser

    const [noMatch, setNoMatch] = useState("")
    const [searchTerm, setSearchTerm] = useState("")
    const [people, setPeople] = useState(null)

    useEffect(() => {
        // ask back end for top 10

        const path = `/getallnewdm/${currentUser.userID}`
        axios.get(`${serverpath}${path}`).then(res => {
            const data = res.data
            if (data.status === "no error") {
                setPeople(data.newDMRes)
            } else {
                console.log(data.status)
            }
        })

    },[])
    const handleKeyPress = (event) => {
        // Check if the Enter key was pressed (key code 13)
        if (event.key === 'Enter') {
            // Trigger the button click action
            searchQuery();
        }
    };

    function searchQuery() {
        if (searchTerm === "") {
            return //maybe error message(?)
        }
        setNoMatch("")
        axios.get(`${serverpath}/search/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(searchTerm)}`).then(res => {
            const data = res.data
            if (data.error === "no error") {
                setPeople(data.userPreviews)
                setNoMatch("")
            } else if (data.error === "no users found") {
                setNoMatch("No Match Found")
                setPeople(null)
            } //catch errors later
        })
    }

    return (
        <div className="flex flex-col justify-start items-center space-y-4">
            <div className="flex flex-start w-full">
                <button className="mb-2 w-fit ml-6 text-3xl hover:text-purple-300" onClick={toggleDMList}> {'←'} </button>
            </div>
            <div className="flex w-11/12 sm:w-3/4 lg:w-1/2 min-w-fit">
                <input 
                    type="text" 
                    value={searchTerm} 
                    onChange={e => setSearchTerm(e.target.value)}
                    onKeyPress={handleKeyPress}
                    className="w-full p-2 rounded-bl-md rounded-tl-md text-black border-1 border-white focus:outline-none focus:border-gray-500 focus:ring-0"
                ></input>
                <div className="relative inset-y-0 right-0 flex items-center px-3 bg-white rounded-tr-md rounded-br-md" onClick={searchQuery}>
                    <img
                        src="./search.png"
                        alt="search users"
                        className="w-6 h-6"
                    ></img>
                </div>
            </div>
            {(noMatch === "No Match Found") ? 
                <div className="w-fit bg-white rounded-lg text-black text-center text-xl mx-auto p-10">
                    {noMatch}
                </div> 
                : null
            }
            {people ? people.map(
                (person, index) => (
                    <Person
                        userID = {person.user_id}
                        back = {toggleSearchUser}
                        showOtherProfile = {toggleOtherProfile}
                        key={index}
                        toggleDMMessage = {toggleDMMessage}
                        username={person.user_name} 
                        user_pic_url={serverpath + person.profile_picture_path}
                    ></Person> 

                )
            ) : null}
        </div>
    );

}

export default NewDM;