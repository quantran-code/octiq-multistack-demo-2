# frozen_string_literal: true

# Keeps sent notifications in memory.
class NotificationStore
  def initialize
    @items = []
  end

  def all
    @items.dup
  end

  def add(recipient, message)
    raise ArgumentError, 'recipient is required' if recipient.to_s.strip.empty?
    raise ArgumentError, 'message is required' if message.to_s.strip.empty?

    item = { id: @items.size + 1, recipient: recipient, message: message }
    @items << item
    item
  end
end
