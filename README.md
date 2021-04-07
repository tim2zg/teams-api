# teams-api

An unofficial wrapper for the unofficial Microsoft Teams API

## Description

This library (still work in progress) was created with the goal
of helping the creation of alternative (and open source) 
Microsoft Teams clients. Currently, the library itself can only
handle a couple of endpoints and is by far not ready yet.  
My goal is to complete this library as soon as possible and start
the creation of an alternative client based on it.

## Usage

For now, you can only run the tests or use the library on your own.
For this you'll need a couple of Microsoft Teams tokens.

### Obtaining a token

Using [teams-token](https://github.com/fossteams/teams-token) one can obtain (and automatically save)
the tokens necessary for this library to work.  
Simply clone and `yarn start` that repository to get your Teams tokens stored into:
```bash
/home/user/.config/fossteams/token-chatsvcagg.jwt
/home/user/.config/fossteams/token-skype.jwt
/home/user/.config/fossteams/token-teams.jwt
```

With these tokens, you'll be able to test out some features like the
`GetConversations` test that retrieves a list of Teams your user is connected with.