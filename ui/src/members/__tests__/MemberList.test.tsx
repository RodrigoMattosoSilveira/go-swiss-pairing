import { render, screen } from '@testing-library/react';
import React from 'react';
import { MemoryRouter } from 'react-router-dom';
import MemberList from '../MemberList';
import { MOCK_MEMBERS } from '../MockMembers';
import userEvent from '@testing-library/user-event';
// import { Provider } from 'react-redux';
// import { store } from '../../state';

describe('<MemberList />', () => {
    const onSave = jest.fn();
    const setup = () => {
        render(
            // <Provider store={store}>
            <MemoryRouter>
                <MemberList members={MOCK_MEMBERS} onSave={onSave}/>
            </MemoryRouter>
            // </Provider>
        );

    }
    beforeEach(() => {});

    test('should render without crashing', () => {
        setup();

        expect(screen).toBeDefined();
    });
    test('should display list', () => {
        setup();
        expect(screen.getAllByRole('button')).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.first' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.last' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.email' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.cell' })).toHaveLength(MOCK_MEMBERS.length);
        expect(screen.getAllByRole('generic', { name: 'member.rating' })).toHaveLength(MOCK_MEMBERS.length);
    });
    test('should display form when edit clicked', async () => {
       setup();
       // eslint-disable-next-line testing-library/render-result-naming-convention
        const user = userEvent.setup();
       await user.click(
             screen.getByRole('button', {name:"m5lpUgrpH"})
           );
       expect(
             screen.getByRole('form', {
                   name: /edit a member/i,
             })
       ).toBeInTheDocument();
    });
    test('should remove form when cancel clicked and show record ', async () => {
        let memeberId = "m5lpUgrpH";
        setup();
        // eslint-disable-next-line testing-library/render-result-naming-convention
        const user = userEvent.setup();
        await user.click(
            screen.getByRole('button', { name: memeberId})
        );
        await user.click(
            screen.getByRole('button', {
                name: /Cancel Edit a Member/i,
            })
        );
        // expect(screen.getByRole('button', {name: memeberId})).toBeInTheDocument();

        // This does not work: https://vijayt.com/post/functional-testing-using-react-testing-library-and-jest/#:~:text=We%20have%20seen%20the%20getByRole%20function.%20It%20retrieves,the%20element%20is%20not%20rendered%20in%20the%20DOM.
        // expect(
        //     screen.getByRole('form', {
        //         name: "Edit a Member",
        //     })
        // ).not.toBeInTheDocument();
        const textElement = screen.queryByText(`Editing member: ${memeberId}`);
        expect(textElement).not.toBeInTheDocument();
    });
})