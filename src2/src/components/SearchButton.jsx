import React from 'react'
import Swal from 'sweetalert2';
function SearchButton({linkValue, isError, setResultResponse}) {
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
    <button 
        className="flex bg-[#68649c] text-center border-black border-[3px] px-8 py-4 rounded-full font-gingerCat text-white shadowButton"
        onClick={handleIsLinkExist}
    >
        SEARCH
    </button>
  )
}

export default SearchButton