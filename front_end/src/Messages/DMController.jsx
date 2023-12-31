import { React, useState, useEffect} from "react";
import DMList from "./DMList";
import DMMessage from "./DMMessage";
import NewDM from "./NewDM"
import OtherProfile from "../Profile/OtherProfile";

function DMController(props) {
    const toggleHomepage = props.toggleHomepage
    const [DMstate, setDMstate] = useState(null);

    // whether to open DM list or messages
    useEffect(() => {
        if (props.wormhole === true) {
            toggleDMMessage(props.friendID)
        } else {
            toggleDMList()
        }
    }, [])

    function toggleDMList(){
        setDMstate(<DMList toggleHomepage={toggleHomepage} toggleDMMessage={toggleDMMessage} toggleNewDM = {toggleNewDM}/>)
    }
    function toggleDMMessage(friendID, friendUsername){
        setDMstate(<DMMessage friendID={friendID} friendUsername={friendUsername} toggleDMList={toggleDMList} toggleOtherProfile={toggleOtherProfile} toggleDMMessage={toggleDMMessage}/>)
    }
    function toggleOtherProfile(userID,backEvent){
        setDMstate(<OtherProfile userID = {userID} goBackScreen = {backEvent}/>)
    }
    function toggleNewDM(friendID){
        setDMstate(<NewDM friendID={friendID} toggleDMMessage={toggleDMMessage} toggleDMList = {toggleDMList} toggleOtherProfile = {toggleOtherProfile}/>)
    }
    return (<div>{DMstate}{/* */}</div>);
}

export default DMController;