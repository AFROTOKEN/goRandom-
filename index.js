const generateACCOUNT =  async( name, email)=>{

    const  sendEmail = await Moralis.Cloud.sendEmail({
        to: recepient
        subject: ${2;subject},
        html: body
    });
}