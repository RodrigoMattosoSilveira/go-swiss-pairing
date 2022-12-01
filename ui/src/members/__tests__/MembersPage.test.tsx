import React from "react";
import { MemoryRouter } from "react-router-dom";
import MembertsPage from "../MembersPage";
import {
    render,
    screen,
    waitForElementToBeRemoved,
} from "@testing-library/react";
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import { url as membersUrl } from '../memberAPI';
import { MOCK_MEMBERS } from '../MockMembers';
// declare which API requests to mock
const server = setupServer(
  // capture "GET http://localhost:3000/projects" requests
      rest.get(membersUrl, (req, res, ctx) => {
          // respond using a mocked JSON body
          return res(ctx.json(MOCK_MEMBERS));
      })
);
beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());
describe("<MembertsPage />", () => {
    function renderComponent() {
        render(
            <MemoryRouter>
                <MembertsPage />
            </MemoryRouter>
        );
    }
    test("should render without crashing", () => {
        renderComponent();
        expect(screen).toBeDefined();
    });
    test('should display loading', () => {
        renderComponent();
        expect(screen.getByText(/loading/i)).toBeInTheDocument();
    });
    test('should display projects', async () => {
        renderComponent();
        expect(await screen.findAllByRole('link')).toHaveLength(MOCK_MEMBERS.length);
    });
    test('should display more button', async () => {
          renderComponent();
          expect(
                await screen.findByRole('button', { name: /more/i })
              ).toBeInTheDocument();
        });
    // this tests the same as the last test but demonstrates
    // what find* methods are doing
    test('should display more button with get', async () => {
        renderComponent();
        await waitForElementToBeRemoved(() => screen.queryByText(/loading/i));
        expect(screen.getByRole('button', { name: /more/i })).toBeInTheDocument();
    });
    test('should display custom error on server error', async () => {
          server.use(
                rest.get(membersUrl, (req, res, ctx) => {
                    return res(ctx.status(500, 'Server error'));
                })
              );
          renderComponent();
              expect(
                  await screen.findByText(/There was an error retrieving the member(s)./i)
              ).toBeInTheDocument();
        });
});