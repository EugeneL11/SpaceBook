import { React, useState } from "react";

function OrbitRequest(props) {

    const acceptRequestEvent = () => {props.acceptRequest(props.username)}
    const denyRequestEvent = () => {props.denyRequest(props.username)}

    return (
        <div class="flex-row">
            <div class="inline-block m-1">
                <img
                    class="rounded-lg w-20 h-20 sr-only"
                    src={props.user_pic_url}
                    alt={props.username} 
                ></img>
                <p>{props.username}</p>
            </div>
            <div class="accept-request-btn inline-block m-1" onClick={acceptRequestEvent}>
                <img 
                    class="sr-only"
                    src={ACCEPT_REQUEST_IMG_URL}
                    alt="accept request"
                ></img>
                <p>accept</p>
            </div>
            <div class="deny-request-btn inline-block m-1" onClick={denyRequestEvent}>
                <img
                    class="sr-only"
                    src={DENY_REQUEST_IMG_URL}
                    alt="deny request"
                ></img>
                <p>deny</p>
            </div>
        </div>
    );
}

function OrbitRequests(props) {
    const [requests, setRequests] = useState(null)

    useEffect(() =>{
        const requeststest = [{
        username: "Gene",
        user_pic_url: "/assets/user0-pfp.jpg"
    }] // placeholder for back-end data

        // ask back-end for request list
        setRequests(requeststest)
    },[])

    const acceptRequest = (requestToAccept) => {
        const newRequestList = requests.filter(
            (request) => request.username !== requestToAccept
        );
        setRequests(newRequestList);
        // do back-end stuff
    }

    const denyRequest = (requestToDeny) => {
        const newRequestList = requests.filter(
            (request) => request.username !== requestToDeny
        );
        setRequests(newRequestList);
        // do back-end stuff
    }
    
    return (
        <div class="flex-col">
            {requests ? requests.map(
                (request, index) => (
                    <div>
                        <OrbitRequest 
                            key={index}
                            username={request.username} 
                            user_pic_url={request.user_pic_url}
                            acceptrequest = {acceptRequest}
                            denyRequest = {denyrequest}
                        ></OrbitRequest>
                    </div>
                )
            ) : null}
        </div>
    );
}

export default OrbitRequests;