# frozen_string_literal: true

require 'rack/test'
require_relative '../lib/app'

RSpec.describe NotifyApp do
  include Rack::Test::Methods

  def app
    NotifyApp
  end

  before { NotifyApp.set :store, NotificationStore.new }

  it 'reports health' do
    get '/health'
    expect(last_response.status).to eq(200)
    expect(JSON.parse(last_response.body)).to eq('status' => 'ok')
  end

  it 'stores a notification and lists it' do
    post '/notifications', JSON.generate(recipient: 'ann', message: 'shipped')
    expect(last_response.status).to eq(201)

    get '/notifications'
    expect(JSON.parse(last_response.body)['notifications'].size).to eq(1)
  end

  it 'rejects a notification without a message' do
    post '/notifications', JSON.generate(recipient: 'ann')
    expect(last_response.status).to eq(400)
    expect(JSON.parse(last_response.body)).to eq('error' => 'message is required')
  end

  it 'rejects a body that is not JSON' do
    post '/notifications', 'nope'
    expect(last_response.status).to eq(400)
  end
end
