import React, { useEffect, useState } from "react";

const Home = (props:{username:string}) => {

  return <div> {props.username?"Hi "+props.username:"you are not logged in"}</div>;
};

export default Home;
