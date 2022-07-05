import React, { useState } from "react";
import styled from "styled-components";
import { Link } from 'react-router-dom';
import { TbHome, TbShoppingCartPlus, Tb3DCubeSphere, TbPlaylistAdd } from 'react-icons/tb';

const FooterContainer = styled.div`
  position: fixed;
  bottom: 0;
  left: 0;
  display: flex;
  justify-content: space-between;
  min-width: 320px;
  width: 100%;
  max-width: 500px;
  padding: 7px 24px;
  background-color: white;

  @media (min-width: 500px) {
    left: calc(50% - 250px);
  }

  &:before {
    content: '';
    position: absolute;
    top: 0;
    left: 8px;
    width: calc(100% - 16px);
    border-top: 1px solid #A0694B;
  }

  a {
    flex: 1 1 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 25%;
    text-decoration: none;
    color: #626262;

    &.active {
      color: #A0694B;
    }

    .footer_icon {
      stroke-width: 1;
    }
  }
`;

const FooterLable = styled.p`

`;

export const Footer = () => {
  
  const [activePage, setActivePage] = useState('home');

  const handleClick = (e) => {
    setActivePage(e.currentTarget.id);
  };

  return(
    <FooterContainer>
      <Link to="/" id="home" onClick={(e) => handleClick(e)} className={activePage === 'home' ? 'active' : null}>
        <TbHome size={24} className="footer_icon"/>
        <FooterLable>Главная</FooterLable>
      </Link>
      <Link to="/stock" id="stock" onClick={(e) => handleClick(e)} className={activePage === 'stock' ? 'active' : null}>
        <Tb3DCubeSphere size={24} className="footer_icon"/>
        <FooterLable>Склад</FooterLable>
      </Link>
      <Link to="/purchase" id="purchase" onClick={(e) => handleClick(e)} className={activePage === 'purchase' ? 'active' : null}>
        <TbShoppingCartPlus size={24} className="footer_icon"/>
        <FooterLable>Закупка</FooterLable>
      </Link>
      <Link to="/expenses" id="expenses" onClick={(e) => handleClick(e)} className={activePage === 'expenses' ? 'active' : null}>
        <TbPlaylistAdd size={24} className="footer_icon"/>
        <FooterLable>Доп.траты</FooterLable>
      </Link>
    </FooterContainer>
  )
}