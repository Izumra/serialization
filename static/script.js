
document.addEventListener("DOMContentLoaded",e=>{
    const petsDOM=document.querySelector(".pets")
    let pets=''
    fetch("http://localhost:6670/pets")
        .then(resp=>resp.json())
        .then(resp=>{
            pets=''
            for(let i=0;i<resp.data?.length;i++){
                pets+=`<div><p>Имя: ${resp.data[i].name}</p><p>Фамилия: ${resp.data[i].age}</p></div>`
            }
            petsDOM.innerHTML=pets
        })
    const addPetButton=document.querySelector(".addPet")
    let submittedPets=[]
    addPetButton.addEventListener("click",(e)=>{
        const name=document.querySelector('input[name="Name"]').value
        const surname=document.querySelector('input[name="Age"]').value
        let data={
            "Name":name,
            "Surname":surname,
        }
        submittedPets.push(data)
        
    })
    const writePetButton=document.querySelector(".writePet")
    writePetButton.addEventListener("click",e=>{
        fetch("http://localhost:6670/pets",{
            headers: {
				'Content-Type': 'application/json'
			},
			method: 'post',
			body: JSON.stringify(submittedPets)
        })
        .then(resp=>resp.json())
        .then(resp=>{
            console.log(resp)
        })
    })
})