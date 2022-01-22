-- Rest model POST: /api/fivem-scripts/auditlog-v1
--[[
    {
        auditId: UUID
        audit: String
        userId: String
        username: String
        type: String
    }
]]
local handleChat = function(author, color, text)
    local auditlog = {
        audit = text,
        userId = author,
        username = GetPlayerName(author),
        type = 'ChatMessage',
    }
    saveAudit(auditlog)
end

local handlePlayerJoining = function(oldId)
    local auditlog = {
        audit = tostring(oldId),
        userId = source,
        username = GetPlayerName(source),
        type = 'playerJoining',
    }
    saveAudit(auditlog)
end

local handlePlayerDropped = function(reason)
    local auditlog = {
        audit = reason,
        userId = source,
        username = GetPlayerName(source),
        type = 'playerDropped',
    }
    saveAudit(auditlog)
end

local handleResponse = function(statusCode, result, resultHeaders)
    if statusCode > 400 then
        print('Error sending auditlogs to server (%i)'):format(statusCode)
    end
    if statusCode ~= 201 then
        print('Unhandled status code (%i)'):format(statusCode)
    end
    local encoded = json.encode(result) -- remove when using actual http endpoint
    local obj = json.decode(tostring(encoded))
    print('Auditlog recorded (' .. obj.type .. ')')
    print('Auditlog recorded (' .. obj.username .. ')')
    print('Auditlog recorded (' .. obj.userId .. ')')
    print('Auditlog recorded (' .. obj.audit .. ')')
end

function saveAudit(auditlog)
    --handleResponse(201, auditlog, nil) -- use for debug
    PerformHttpRequest('http://localhost:8090/api/fivem-scripts/auditlog-v1', handleResponse, "POST", json.encode(auditlog))
end

AddEventHandler('chatMessage', handleChat)
AddEventHandler('playerDropped', handlePlayerDropped)
AddEventHandler('playerJoining', handlePlayerJoining)