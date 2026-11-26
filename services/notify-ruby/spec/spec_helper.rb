# frozen_string_literal: true

require 'simplecov'
require 'simplecov_json_formatter'

SimpleCov.formatter = SimpleCov::Formatter::JSONFormatter
SimpleCov.start do
  add_filter '/spec/'
  add_filter '/vendor/'
end

ENV['APP_ENV'] = 'test'
ENV['RACK_ENV'] = 'test'
