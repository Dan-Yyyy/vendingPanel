import React from "react";
import styled from "styled-components";
import { useSelector, useDispatch } from 'react-redux';

import { TitleBlock } from '../components/TitleBlock';
import { Point } from "../components/Point";
import { FiEdit } from 'react-icons/fi';
import { setActiveDay } from "../redux/home/homeReduser";


const HomeContainer = styled.div`
  padding-bottom: 58px;
`;

const TaskBlock = styled.div`
  margin-bottom: 48px;
`;

const TaskMenu = styled.ul`
  position: relative;
  display: flex;
  justify-content: space-between;
`;

const TaskTime = styled.li`
  color: #626262;
  font-size: 18px;
  font-weight: 300;
  font-family: "Roboto-Thin";
  list-style: none;
  cursor: pointer;

  &.active {
    color: #000;
    font-family: "Roboto-Regular";
  }
`;

const DotItem = styled.div`
    position: absolute;
    bottom: -8px;
    width: 7px;
    height: 7px;
    border-radius: 50%;
    background-color: black;
    transition: all 0.4s;

    ${ props => props.active === 'today' && `
      left: 3em;
    `}
    ${ props => props.active === 'tomorrow' && `
      left: 50%;
    `}
    ${ props => props.active === 'week' && `
      left: calc(100% - 3em);
    `}
`;

const PointContainer = styled.div``;


export const Home = () => {
  const home = useSelector(state => state.home);
  const dispatch = useDispatch();
  const active = home.activeDate;

  const handleClick = (e) => {
    if(e.target.id !== active)
      dispatch(setActiveDay(`${e.target.id}`));
  }

  return(
    <HomeContainer>
      <TitleBlock title={ 'План объезда' } Icon={ FiEdit }/>
      <TaskBlock>
        <TaskMenu>
          <TaskTime id='today' onClick={(e) => handleClick(e)} className={active === 'today' ? 'active' : null}>На сегодня</TaskTime>
          <TaskTime id='tomorrow' onClick={(e) => handleClick(e)} className={active === 'tomorrow' ? 'active' : null}>На завтра</TaskTime>
          <TaskTime id='week' onClick={(e) => handleClick(e)} className={active === 'week' ? 'active' : null}>На неделю</TaskTime>
          <DotItem active={active}/>
        </TaskMenu>
      </TaskBlock>
      <PointContainer>
        {
          home.status === 'fullfield'
          ? home.items.map(point => 
            <Point adress={point.adress} 
              mashine={point.mashine} 
              isVisited={point.isVisited} 
              key={point.id + point.adress} 
              id={point.id}/>)
          : null
        }     
      </PointContainer>
    </HomeContainer>
  )
};