const msg = require('../demo');

test('The message: Hello World! Should be returned.', () => {
    expect(msg.hello()).toBe('Hello World!');
});