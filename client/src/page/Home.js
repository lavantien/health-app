import { useEffect, useState } from 'react';

const Home = () => {
	const [data, setData] = useState([]);

	useEffect(() => {
		const fetchData = async () => {
			const res = await fetch('http://localhost:8080/users');
			const data = await res.json();
			console.log(data);
			setData(data);
		};
		fetchData();
	}, []);

	return (
		<>
			<h1 className='blue'>Homepage</h1>

			<h2>Users List:</h2>
			<ol>
				{data ? (
					data.map((item, index) => (
						<li key={index}>
							<span>{item.id}</span> <span>{item.name}</span>
						</li>
					))
				) : (
					<li>Loading...</li>
				)}
			</ol>
		</>
	);
};

export default Home;
