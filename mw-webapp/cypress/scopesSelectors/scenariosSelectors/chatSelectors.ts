import {chatAccessIds} from "cypress/accessIds/scenariosAccessIds/chatAccessIds";
import {getDataCy} from "src/utils/cyTesting/getDataCy";

export const chatSelectors = {
    getOpenChatButton: () => cy.get(getDataCy(chatAccessIds.openChatButton)),
    getmessagesAmount: () => cy.get(getDataCy(chatAccessIds.messagesAmount)),

    chatContainer:{
        getChatItem: (name: string) => cy.get(getDataCy(chatAccessIds.chatContainer.chatItem(name))),
        getChatItemName: () => cy.get(getDataCy(chatAccessIds.chatContainer.chatItemName)),
        getMessageInput: () => cy.get(getDataCy(chatAccessIds.chatContainer.messageInput)),
        getSendMessageButton: () => cy.get(getDataCy(chatAccessIds.chatContainer.sendMessageButton)),
        getMessageItem: () => cy.get(getDataCy(chatAccessIds.chatContainer.messageItem))
    }
};