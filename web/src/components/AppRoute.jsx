import React from "react";
import { Routes, Route } from 'react-router-dom';
import styled from "styled-components";

import { pablicRoutes, privateRoutes } from './../router';
import { Header } from './Header';
import { Footer } from "./Footer";
import { useSelector } from "react-redux";

const PageContainer = styled.div`
  position: relative;
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 24px;
`;

export const AppRoute = () => {
  const isAuth = useSelector(state => state.auth.isAuth);

  return(
    <>
      {
        isAuth
        ? 
          <PageContainer>
            <Header />
            <Routes>
              {
                privateRoutes.map(route => <Route path={route.path} exact={route.exact} element={<route.element/>} key={route.path}/>)
              }
            </Routes>
            <Footer/>
          </PageContainer>
        :
          <Routes>
            {
              pablicRoutes.map(route => <Route path={route.path} exact={route.exact} element={<route.element/>} key={route.path}/>)
            }
          </Routes>
      }
    </>
  )
}