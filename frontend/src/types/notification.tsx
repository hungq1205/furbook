export interface Notification {
    id: string
    username: string
	icon: string
	desc: string
    link: string
	read: boolean
    created_at: string
}

export const formatNotification = (noti: Notification): Notification => {
    return {
        id: noti.id,
        username: noti.username,
        icon: getIcon(noti.icon),
        desc: getDescription(noti),
        link: getLink(noti),
        read: noti.read,
        created_at: noti.created_at,
    } as Notification;
} 

const getLink = (noti: Notification): string => {
    const descParts = noti.desc.split(":")
    switch (descParts[0]) {
        case "post": return `post/${noti.link}`;
        case "friendRequest": return `profile/${noti.link}`;
    }
    return noti.desc
}

const getDescription = (noti: Notification): string => {
    const descParts = noti.desc.split(":")
    switch (descParts[0]) {
        case "post": return getPostDescription(noti);
        case "friendRequest": return getFriendRequestDescription(noti);
    }
    return noti.desc
}

const getPostDescription = (noti: Notification): string => {
    const descParts = noti.desc.split(":")
    switch (descParts[1]) {
        case "resolved": return `${descParts[2]} has resolved a lost/found`
        case "comment": return `${descParts[2]} commented on your post`;
        case "interaction": return `${descParts[2]} interacted with your post`;
        case "participate": return `${descParts[2]} participated in the your lost/found post`;
        case "unparticipate": return `${descParts[2]} unparticipated from the your lost/found post`;
    }
    return noti.desc
}

const getFriendRequestDescription = (noti: Notification): string => {
    const descParts = noti.desc.split(":")
    switch (descParts[1]) {
        case "send": return `${noti.link} sent you a friend request`;
        case "accepted": return `${noti.link} accepted your friend request`;
        case "declined": return `${noti.link} declined your friend request`;
    }
    return noti.desc
}

const getIcon = (icon: string): string => {
    switch (icon) {
        case "post": return "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSFG1_ZigPGI8JWWzb4KPJkqHWc8LisbBRZbg&s"
        case "comment": return "https://cdn.iconscout.com/icon/free/png-256/free-comment-icon-download-in-svg-png-gif-file-formats--chat-message-communication-conversation-user-interface-vol-3-pack-icons-2202811.png"
        case "interaction": return  "https://cdn4.iconfinder.com/data/icons/web-ui-ux/32/007-Heart-512.png"
    }
    return icon
}