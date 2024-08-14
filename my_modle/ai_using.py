import ai_modle

def AI_MODLE(option='QW_MAX', user_input='你好', path='', promot='你是一个有帮助的AI助手'):      # loading
    if option == 'QW-Max':
        return ai_modle.QW_MAX(user_input, path, promot)
    # Insert into this line
