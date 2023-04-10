// import { BrowserRouter, Outlet, Route, Routes } from 'react-router-dom';
// import { Outlet, RouterProvider, createBrowserRouter } from 'react-router-dom';
import { Outlet, RouterProvider, createHashRouter } from 'react-router-dom';
import About from '../page/About';
import Home from '../page/Home';
import Footer from './Footer';
import Header from './Header';

const Router = () => {
	const Layout = () => {
		return (
			<>
				<Header />
				<Outlet />
				<Footer />
			</>
		);
	};

	// const BrowserRoutes = () => {
	// 	return (
	// 		<>
	// 			<BrowserRouter>
	// 				<Routes>
	// 					<Route path='/' element={<Layout />}>
	// 						<Route path='/' element={<Home />} />
	// 						<Route path='/about' element={<About />} />
	// 					</Route>
	// 				</Routes>
	// 			</BrowserRouter>
	// 		</>
	// 	);
	// };

	const BrowserRoutes = createHashRouter([
		{
			path: '/',
			element: <Layout />,
			children: [
				{
					path: '/',
					element: <Home />,
				},
				{
					path: '/about',
					element: <About />,
				},
			],
		},
	]);

	return (
		<>
			{/* <BrowserRoutes /> */}
			<RouterProvider router={BrowserRoutes} />
		</>
	);
};

export default Router;
