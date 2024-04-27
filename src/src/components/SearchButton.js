import { link } from "d3";
import React from "react";
import Swal from "sweetalert2";
function SearchButton({isOn,linkValue, setResultResponse, isName, isError}) {
    const showSuccess = () => {
        Swal.fire({
          position: "center",
          icon: "success",
          title: "Link Valid! Start Searching  ",
          showConfirmButton: false,
          timer: 1500
        });
        
    }

    const handleIsLinkExist = async () => {
        try {
            let whatError = []
            for (const key in linkValue) {
                let response = await fetch(
                    `https://api.allorigins.win/get?url=${encodeURIComponent(
                        `https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=${linkValue[key]}`
                    )}`
                )

                const data = await response.json()
                const jsonData = JSON.parse(data.contents)
                // console.log(jsonData.query.search.map(item => item.title))
                if (jsonData.query.search.map(item => item.title).length === 0){
                    whatError.push(linkValue[key])
                }
            }
            isError(whatError)
            if (whatError.length === 0){
                showSuccess()
                handleClick()
            }
        } catch (error) {
            console.error('Error fetching suggestions:', error);
        }
    }

    const handleClick = async () => {
        try {
            const response = await fetch('http://localhost:8000', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    linkValue: linkValue.startLink,
                    finValue: linkValue.endLink,
                    isOn: isOn,
                    isName: isName,
                }),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const tempResponse = await response.json();
            setResultResponse({
                exec: tempResponse.exec,
                len: tempResponse.len,
                urls: tempResponse.url,
                resultLink: tempResponse.links,
            })
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <button onClick={handleIsLinkExist} style={{fontFamily:"Poppins",fontSize:"20px",border:"2px solid white",borderRadius:"10px",backgroundColor:"black",color:"white",width:"150px",height:"40px", marginTop: '20px'}}>Start Racing</button>
    );
}

export default SearchButton;