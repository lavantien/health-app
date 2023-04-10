import reactIcon from '../assets/image/react-icon.png';

const Header = () => {
	return (
		<nav className='nav-bar'>
			<img src={reactIcon} alt='react-icon' height='40' />
			<ul>
				<li>
					<a href='/'>Home</a>
				</li>
				<li>
					<a href='/about'>About</a>
				</li>
			</ul>
		</nav>
	);
};

export default Header;
