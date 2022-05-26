export const getProjectsOfUser =  async (id:any) => {
    let res = await fetch(`http://localhost:1234/projects/${id}`, {
        
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
    });
    const content = await res.json();
    console.log('Content: ',content)
    return content

    // let array = []
    // array.push({
    //     ID: 10,
    //     CreatedAt: "2022-05-26T21:34:18.144414+03:00",
    //     UpdatedAt: "2022-05-26T22:38:07.510554+03:00",
    //     DeletedAt: null,
    //     user_id: 1,
    //     project_name: "asdqwe",
    //     info: "zxczxczxc",
    //     size: 185075623,
    //     link: "http://localhost:1234/projects/1/kekich.html"
    // })
    // return array

}