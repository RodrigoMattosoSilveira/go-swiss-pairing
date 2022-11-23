import React from 'react';
import App from './App';
import renderer from "react-test-renderer";

describe('<App />', () => {
  // test('renders home heading', () => {
  //     // render(<HomePage/>);
  //     // expect(screen.getByRole('heading')).toHaveTextContent('Home');
  //     test('should render without crashing', () => {
  //         render(<HomePage />);
  //     });
  // });
  test('snapshot', () => {
    const tree = renderer.create(<App />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
