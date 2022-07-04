import { Home } from './../pages/Home';
import { Login } from './../pages/Login';
import { Stock } from './../pages/Stock';
import { Purchase } from './../pages/Purchase';
import { Expenses } from './../pages/Expenses';

const RouteName = {
  HOME: '/',
  LOGIN: '/login',
  DEFAULT: '*',
  STOCK: '/stock',
  PURCHASE: '/purchase',
  EXPENSES: '/expenses'
};

export const pablicRoutes = [
  {
    path: RouteName.LOGIN, exact: true, element: Login
  },
  {
    path: RouteName.DEFAULT, exact: true, element: Login
  }
];

export const privateRoutes = [
  {
    path: RouteName.HOME, exact: true, element: Home
  },
  {
    path: RouteName.DEFAULT, exact: true, element: Home
  },
  {
    path: RouteName.STOCK, exact: true, element: Stock
  },
  {
    path: RouteName.PURCHASE, exact: true, element: Purchase
  },
  {
    path: RouteName.EXPENSES, exact: true, element: Expenses
  }
]