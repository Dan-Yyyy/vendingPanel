import styled from 'styled-components';

const TitleContainer = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 32px;
  
  .icon {
    color: #A0694B;
    cursor: pointer;
    stroke-width: 1;
  }
`;

const Title = styled.h1`
  color: #000;
  font-size: 32px;
  font-weight: 400;
  text-transform: uppercase;
`;

export const TitleBlock = ({title, Icon}) => {
  return(
    <TitleContainer>
      <Title>{title}</Title>
      <Icon size={36} className="icon" />
    </TitleContainer>
  )
}
