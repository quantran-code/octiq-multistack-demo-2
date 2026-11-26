# frozen_string_literal: true

require 'json'
require 'sinatra/base'
require_relative 'notification_store'

# HTTP API for the notification service.
class NotifyApp < Sinatra::Base
  set :store, NotificationStore.new

  before { content_type :json }

  get '/health' do
    JSON.generate(status: 'ok')
  end

  get '/notifications' do
    JSON.generate(notifications: settings.store.all)
  end

  post '/notifications' do
    payload = JSON.parse(request.body.read)
    halt 400, JSON.generate(error: 'body must be a JSON object') unless payload.is_a?(Hash)

    status 201
    JSON.generate(settings.store.add(payload['recipient'], payload['message']))
  rescue JSON::ParserError
    halt 400, JSON.generate(error: 'body must be valid JSON')
  rescue ArgumentError => e
    halt 400, JSON.generate(error: e.message)
  end
end
